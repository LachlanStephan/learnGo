DROP TABLE IF EXISTS Users;
CREATE TABLE Users (
    User_id int NOT NULL AUTO_INCREMENT,
    FirstName varchar(255),
    LastName varchar(255),
    Password varchar(255),
    IsAdmin tinyint,
    Created_at timestamp,
    Updated_at timestamp,
    PRIMARY KEY(User_id)
) ENGINE=INNODB;

DROP TABLE IF EXISTS Blogs;
CREATE TABLE Blogs (
    Blog_id int NOT NULl AUTO_INCREMENT,
    User_id int,
    Title varchar(255),
    Content longtext,
    Created_at timestamp,
    Updated_at timestamp,
    PRIMARY KEY(Blog_id),
    FOREIGN KEY(User_id) REFERENCES Users(User_id)
) ENGINE=INNODB;

DROP TABLE IF EXISTS Tags;
CREATE TABLE Tags (
    Tag_id int NOT NULL AUTO_INCREMENT,
    Tag_name varchar(255),
    Hex_code varchar(255),
    Description varchar(255),
    Created_at timestamp,
    Updated_at timestamp,
    PRIMARY KEY(Tag_id)
) ENGINE=INNODB;

DROP TABLE IF EXISTS Blogs_tags;
CREATE TABLE Blogs_tags (
    Blogs_tags_id int NOT NULL AUTO_INCREMENT,
    Blog_id int,
    Tag_id int,
    PRIMARY KEY(Blogs_tags_id),
    FOREIGN KEY(Blog_id) REFERENCES Blogs(Blog_id),
    FOREIGN KEY(Tag_id) REFERENCES Tags(Tag_id)
) ENGINE=INNODB;
