-- +goose Up
CREATE TABLE match_result (
  id            INTEGER PRIMARY KEY,
  matchid       INT,
  playnumber    INT,
  matchtype     VARCHAR(16),
  redscorejson  TEXT,
  bluescorejson TEXT,
  redcardsjson  TEXT,
  bluecardsjson TEXT
);

-- +goose Down
DROP TABLE match_results;
