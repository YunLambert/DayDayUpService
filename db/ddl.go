package db

const (
	createDailySignTable = `
CREATE TABLE IF NOT EXISTS daily_sign
(
    id         INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id    TEXT    NOT NULL UNIQUE ,
    january    INTEGER NOT NULL DEFAULT 0,
    february   INTEGER NOT NULL DEFAULT 0,
    march      INTEGER NOT NULL DEFAULT 0,
    april      INTEGER NOT NULL DEFAULT 0,
    may        INTEGER NOT NULL DEFAULT 0,
    june       INTEGER NOT NULL DEFAULT 0,
    july       INTEGER NOT NULL DEFAULT 0,
    august     INTEGER NOT NULL DEFAULT 0,
    september  INTEGER NOT NULL DEFAULT 0,
    october    INTEGER NOT NULL DEFAULT 0,
    november   INTEGER NOT NULL DEFAULT 0,
    december   INTEGER NOT NULL DEFAULT 0,
    streak     INTEGER NOT NULL DEFAULT 0,
    sign_today INTEGER NOT NULL DEFAULT 0
);
`
)
