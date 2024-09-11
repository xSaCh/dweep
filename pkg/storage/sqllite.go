package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xSaCh/dweep/pkg/models"
)

type SqlliteStore struct {
	db *sql.DB
}

func NewSqlliteStore(file string) (*SqlliteStore, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}
	return &SqlliteStore{db: db}, nil
}

func (s *SqlliteStore) Create() error {

	queryW := `CREATE TABLE WatchlistItem (
		WatchlistItemId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		UserID INTEGER NOT NULL,
		FilmId INTEGER NOT NULL,
		Type char(20),
		MyRating REAL,
		WatchStatus char(20) NOT NULL,
		Note TEXT,
		AddedOn DATE,
		UpdatedOn DATE
		);`
	// PRIMARY KEY (WatchlistItemId, UserID),

	queryR := `CREATE TABLE WatchlistItem_Recommended ( WatchlistItemId INTEGER, userId INTEGER, RecommendedBy INTEGER,
		FOREIGN KEY (WatchlistItemId) REFERENCES WatchlistItem(WatchlistItemId));`
	queryT := `CREATE TABLE WatchlistItem_Tag ( WatchlistItemId INTEGER, userId INTEGER, Tag INTEGER,
		FOREIGN KEY (WatchlistItemId) REFERENCES WatchlistItem(WatchlistItemId));`

	queryM := `CREATE TABLE WatchlistItem_Movie ( WatchlistItemId INTEGER, userId INTEGER, watchedDate DATE,
		FOREIGN KEY (WatchlistItemId) REFERENCES WatchlistItem(WatchlistItemId));`

	s.db.Exec("PRAGMA foreign_keys = ON;")
	s.db.Exec("DROP TABLE IF EXISTS WatchlistItem_Recommended;")
	s.db.Exec("DROP TABLE IF EXISTS WatchlistItem_Tag;")
	s.db.Exec("DROP TABLE IF EXISTS WatchlistItem_Movie;")
	s.db.Exec("DROP TABLE IF EXISTS WatchlistItem;")

	_, err := s.db.Exec(queryW)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}
	_, err = s.db.Exec(queryR)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}
	_, err = s.db.Exec(queryT)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}
	_, err = s.db.Exec(queryM)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}
	return nil
}

func (s *SqlliteStore) AddMovie(item models.ReqWatchlistItemMovie, filmId int, userId int) error {
	queryW := `INSERT INTO WatchlistItem (UserId, FilmId, Type, MyRating, WatchStatus, Note, AddedOn, UpdatedOn) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	if s.getWatchlistId(filmId, userId) != -1 {
		return fmt.Errorf("movie already exists in watchlist")
	}

	r, err := s.db.Exec(queryW, userId, filmId, "movie", item.MyRating, item.WatchStatus, item.Note, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("error inserting movie: %v", err)
	}
	wid, _ := r.LastInsertId()
	if err := s.setWatchlistMovieOtherData(item, int(wid), userId); err != nil {
		return err
	}

	return nil
}

func (s *SqlliteStore) UpdateMovie(item models.ReqWatchlistItemMovie, filmId int, userId int) error {
	queryW := `UPDATE WatchlistItem SET MyRating = ?, WatchStatus = ?, Note = ?, UpdatedOn = ? WHERE FilmId = ? AND UserID = ?;`

	queryDM := `DELETE FROM WatchlistItem_Movie WHERE WatchlistItemId = ? AND userId = ?;`
	queryDT := `DELETE FROM WatchlistItem_Tag WHERE WatchlistItemId = ? AND userId = ?;`
	queryDR := `DELETE FROM WatchlistItem_Recommended WHERE WatchlistItemId = ? AND userId = ?;`

	watchlistId := s.getWatchlistId(filmId, userId)
	if watchlistId == -1 {
		return fmt.Errorf("filmId not exists")
	}

	r, err := s.db.Exec(queryW, item.MyRating, item.WatchStatus, item.Note, time.Now(), filmId, userId)
	ra, _ := r.RowsAffected()
	if err != nil && ra != 1 {
		return err
	}
	// Remove all tags, recommendedUserId and watchedDates befre reinserting
	s.db.Exec(queryDM, watchlistId, userId)
	s.db.Exec(queryDT, watchlistId, userId)
	s.db.Exec(queryDR, watchlistId, userId)

	if err := s.setWatchlistMovieOtherData(item, watchlistId, userId); err != nil {
		return err
	}
	return nil
}
func (s *SqlliteStore) RemoveMovie(filmId int, userId int) error {
	queryDW := `DELETE FROM WatchlistItem WHERE WatchlistItemId = ? AND userId = ?;`
	queryDM := `DELETE FROM WatchlistItem_Movie WHERE WatchlistItemId = ? AND userId = ?;`
	queryDT := `DELETE FROM WatchlistItem_Tag WHERE WatchlistItemId = ? AND userId = ?;`
	queryDR := `DELETE FROM WatchlistItem_Recommended WHERE WatchlistItemId = ? AND userId = ?;`

	watchlistId := s.getWatchlistId(filmId, userId)
	if watchlistId == -1 {
		return fmt.Errorf("filmId not exists")
	}
	// Remove all tags, recommendedUserId and watchedDates befre reinserting
	if _, err := s.db.Exec(queryDM, watchlistId, userId); err != nil {
		return fmt.Errorf("error deleting movie watched dates: %v", err)
	}
	if _, err := s.db.Exec(queryDT, watchlistId, userId); err != nil {
		return fmt.Errorf("error deleting movie tags: %v", err)
	}
	if _, err := s.db.Exec(queryDR, watchlistId, userId); err != nil {
		return fmt.Errorf("error deleting movie Recommended by: %v", err)
	}
	if _, err := s.db.Exec(queryDW, watchlistId, userId); err != nil {
		return fmt.Errorf("error deleting movie: %v", err)
	}

	return nil
}
func (s *SqlliteStore) GetAllMovies(userId int) ([]models.WatchlistItemMovie, error) {
	query := `SELECT WatchlistItemId, FilmId, Type, MyRating, WatchStatus, Note, AddedOn, UpdatedOn FROM WatchlistItem WHERE UserID = ?;`

	rows, err := s.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error querying movies: %v", err)
	}
	defer rows.Close()

	var movies []models.WatchlistItemMovie
	for rows.Next() {
		var movie models.WatchlistItemMovie
		err := rows.Scan(&movie.WatchlistItemId, &movie.FilmId, &movie.Type, &movie.MyRating, &movie.WatchStatus, &movie.Note, &movie.AddedOn, &movie.UpdatedOn)
		if err != nil {
			return nil, fmt.Errorf("error scanning movie: %v", err)
		}

		s.getWatchlistMovieOtherData(&movie, userId)
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over movies: %v", err)
	}

	return movies, nil
}
func (s *SqlliteStore) GetMovie(filmId int, userId int) (models.WatchlistItemMovie, error) {
	wid := s.getWatchlistId(filmId, userId)
	if wid == -1 {
		return models.WatchlistItemMovie{}, fmt.Errorf("movie not found")
	}

	query := `SELECT WatchlistItemId, FilmId, Type, MyRating, WatchStatus, Note, AddedOn, UpdatedOn FROM WatchlistItem WHERE WatchlistItemId = ? AND UserID = ?;`
	row := s.db.QueryRow(query, wid, userId)

	var movie models.WatchlistItemMovie
	err := row.Scan(&movie.WatchlistItemId, &movie.FilmId, &movie.Type, &movie.MyRating, &movie.WatchStatus, &movie.Note, &movie.AddedOn, &movie.UpdatedOn)
	if err != nil {
		return models.WatchlistItemMovie{}, fmt.Errorf("error scanning movie: %v", err)
	}

	s.getWatchlistMovieOtherData(&movie, userId)
	return movie, nil
}

func (s *SqlliteStore) getWatchlistMovieOtherData(w *models.WatchlistItemMovie, userId int) {
	queryM := `SELECT watchedDate FROM WatchlistItem_Movie WHERE WatchlistItemId = ? AND userId = ?;`
	queryT := `SELECT Tag FROM WatchlistItem_Tag WHERE WatchlistItemId = ? AND userId = ?;`
	queryR := `SELECT RecommendedBy FROM WatchlistItem_Recommended WHERE WatchlistItemId = ? AND userId = ?;`

	rows, err := s.db.Query(queryM, w.WatchlistItemId, userId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d time.Time
		err := rows.Scan(&d)
		if err != nil {
			return
		}
		w.WatchedDates = append(w.WatchedDates, d)
	}

	rows, err = s.db.Query(queryT, w.WatchlistItemId, userId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			return
		}
		w.MyTags = append(w.MyTags, t)
	}

	rows, err = s.db.Query(queryR, w.WatchlistItemId, userId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r int
		err := rows.Scan(&r)
		if err != nil {
			return
		}
		w.RecommendedBy = append(w.RecommendedBy, int64(r))
	}
}

func (s *SqlliteStore) setWatchlistMovieOtherData(item models.ReqWatchlistItemMovie, watchlistId int, userId int) error {
	queryM := `INSERT INTO WatchlistItem_Movie (WatchlistItemId, userId, watchedDate) VALUES (?, ?, ?);`
	queryT := `INSERT INTO WatchlistItem_Tag (WatchlistItemId, userId, Tag) VALUES (?, ?, ?);`
	queryR := `INSERT INTO WatchlistItem_Recommended (WatchlistItemId, userId, RecommendedBy) VALUES (?, ?, ?);`

	for i := range item.MyTags {
		_, err := s.db.Exec(queryT, watchlistId, userId, item.MyTags[i])
		if err != nil {
			return fmt.Errorf("error inserting movie tags: %v", err)
		}
	}
	for i := range item.RecommendedBy {
		_, err := s.db.Exec(queryR, watchlistId, userId, item.RecommendedBy[i])
		if err != nil {
			return fmt.Errorf("error inserting movie recommended by: %v", err)
		}
	}

	for i := range item.WatchedDates {
		fmt.Printf("[] watched date: %v\n", item.WatchedDates[i])
		_, err := s.db.Exec(queryM, watchlistId, userId, item.WatchedDates[i])
		if err != nil {
			return fmt.Errorf("error inserting movie watched dates: %v", err)
		}
	}

	return nil
}

func (s *SqlliteStore) getWatchlistId(filmid int, userid int) int {
	query := `SELECT WatchlistItemId FROM WatchlistItem WHERE FilmId = ? AND UserID = ?;`
	rows, err := s.db.Query(query, filmid, userid)
	if err != nil || !rows.Next() {
		return -1
	}
	defer rows.Close()
	wid := -1
	rows.Scan(&wid)
	return wid
}
