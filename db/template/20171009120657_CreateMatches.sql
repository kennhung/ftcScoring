-- +goose Up
CREATE TABLE matches (
  id           INTEGER PRIMARY KEY,
  type         VARCHAR(16),
  displayname  VARCHAR(16),
  time         DATETIME,
  red1         INT,
  red1notshow  bool,
  red2         INT,
  red2notshow  bool,
  blue1        INT,
  blue1notshow bool,
  blue2        INT,
  blue2notshow bool,
  status       VARCHAR(16),
  startedat    DATETIME,
  winner       VARCHAR(16)
);

-- +goose Down
DROP TABLE matches;
