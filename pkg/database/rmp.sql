-- name: create-types-table
CREATE TABLE types (
       id_type INTEGER PRIMARY KEY,
       description TEXT
);

-- name: create-type0
INSERT INTO types VALUES(0, 'Person');
-- name: create-type1
INSERT INTO types VALUES(1, 'Group');
-- name: create-type2
INSERT INTO types VALUES(2, 'Unknown');

-- name: create-performers-table
CREATE TABLE performers (
       id_performer INTEGER PRIMARY KEY,
       id_type INTEGER,
       name TEXT,
       FOREIGN KEY (id_type) REFERENCES types(id_type)
);

-- name: create-persons-table
CREATE TABLE persons(
       id_person INTEGER PRIMARY KEY,
       stage_name TEXT,
       real_name TEXT,
       birth_date TEXT,
       death_date TEXT
);

-- name: create-groups-table
CREATE TABLE groups (
       id_group INTEGER PRIMARY KEY ,
       name TEXT ,
       start_date TEXT ,
       end_date TEXT
);

-- name: create-albums-table
CREATE TABLE albums (
       id_album INTEGER PRIMARY KEY ,
       path TEXT,
       name TEXT,
       year INTEGER
);

-- name: create-rolas-table
CREATE TABLE rolas (
       id_rola INTEGER PRIMARY KEY ,
       id_performer INTEGER ,
       id_album INTEGER ,
       path TEXT ,
       title TEXT ,
       track INTEGER ,
       year INTEGER ,
       genre TEXT ,
       FOREIGN KEY ( id_performer ) REFERENCES performers ( id_performer ) ,
       FOREIGN KEY ( id_album ) REFERENCES albums ( id_album )
);

-- name: create-in_group-table
CREATE TABLE in_group (
       id_person INTEGER,
       id_group INTEGER,
       PRIMARY KEY (id_person, id_group),
       FOREIGN KEY (id_person) REFERENCES persons(id_person),
       FOREIGN KEY (id_group) REFERENCES groups(id_group)
);

-- name: insert-rola
INSERT INTO rolas (id_rola, id_performer, id_album, path, title, track, year,
genre) VALUES (?, ?, ?, ?, ?, ?, ?, ?)

-- name: insert-album
INSERT INTO albums (id_album, path, name, year) VALUES (?, ?, ?, ?)

-- name: insert-group
INSERT INTO  groups (id_group, name, start_date, end_date) VALUES (?, ?, ?, ?)

-- name: insert-in_group
INSERT INTO in_groups (id_person, id_group) VALUES (?, ?)

-- name: insert-person
INSERT INTO persons (id_person, stage_name, real_name, birth_date, death_date) VALUES (?, ?, ?, ?, ?)

-- name: insert-performer
INSERT INTO performers (id_performer, id_type, name) VALUES (?, ?, ?)

-- name: find-rolas-by-id
SELECT * FROM rolas WHERE id_rola = ?

