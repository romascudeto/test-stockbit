-- CREATE
CREATE TABLE USER (
  ID INTEGER PRIMARY KEY,
  UserName TEXT NOT NULL,
  Parent INTEGER NULL
);

-- INSERT
INSERT INTO USER VALUES (1, 'Ali', 2);
INSERT INTO USER VALUES (2, 'Budi', 0);
INSERT INTO USER VALUES (3, 'Cecep', 1);

-- QUERY 
SELECT u1.ID, u1.UserName, u2.UserName 
FROM USER u1
LEFT JOIN USER u2
ON u1.Parent = u2.ID
