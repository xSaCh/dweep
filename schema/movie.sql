CREATE TABLE Movies (
    FilmId      INTEGER PRIMARY KEY NOT NULL AUTOINCREMENT,
    TmdbId      INTEGER NOT NULL,
    ImdbId      CHAR(20) NOT NULL,
    Title       TEXT NOT NULL,
    ReleaseDate DATE NOT NULL,
    Runtime     INTEGER,
    AgeRating   TEXT,
    Overview    TEXT,
    PosterUrl   TEXT,
    Rating      REAL,
    Director    TEXT
);

CREATE TABLE Genres (
    FilmId  INTEGER,
    Genre   TEXT,
    PRIMARY KEY (FilmId, Genre),
    FOREIGN KEY (FilmId) REFERENCES Movies (FilmId)
);

CREATE TABLE CastIds (
    FilmId  INTEGER PRIMARY KEY AUTOINCREMENT,
    Name    TEXT NOT NULL UNIQUE
);

CREATE TABLE Casts (
    FilmId  INTEGER,
    CastId  INTEGER,
    PRIMARY KEY (FilmId, CastId),
    FOREIGN KEY (FilmId) REFERENCES Movies (FilmId),
    FOREIGN KEY (CastId) REFERENCES CastIds (FilmId)
);


CREATE TABLE Tags (
    MovieId INTEGER,
    Tag     TEXT,
    FOREIGN KEY (MovieId) REFERENCES Movies (FilmId),
);
