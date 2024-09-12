CREATE TABLE WatchlistItem (
    WatchlistItemId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    UserID          INTEGER NOT NULL,
    FilmId          INTEGER NOT NULL,
    Type            CHAR(20),
    MyRating        REAL,
    WatchStatus     CHAR(20) NOT NULL,
    Note            TEXT,
    AddedOn         DATE,
    UpdatedOn       DATE
);

CREATE TABLE WatchlistItem_Recommended (
    WatchlistItemId INTEGER,
    UserId          INTEGER,
    RecommendedBy   INTEGER,
	FOREIGN KEY (WatchlistItemId) REFERENCES WatchlistItem(WatchlistItemId)
);

CREATE TABLE WatchlistItem_Tag ( 
    WatchlistItemId INTEGER,
    UserId          INTEGER,
    Tag             TEXT,
    FOREIGN KEY (WatchlistItemId) REFERENCES WatchlistItem(WatchlistItemId)
);

CREATE TABLE WatchlistItem_Movie ( 
    WatchlistItemId INTEGER,
    UserId          INTEGER,
    watchedDate     DATE,
    FOREIGN KEY (WatchlistItemId) REFERENCES WatchlistItem(WatchlistItemId)
);