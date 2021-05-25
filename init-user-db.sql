CREATE TABLE users
(
    email character varying(255) PRIMARY KEY,
    hashpassword character varying(255) NOT NULL,
    firstname character varying(255),
    lastname character varying(255)
);
ALTER TABLE users
    OWNER to postgres;
