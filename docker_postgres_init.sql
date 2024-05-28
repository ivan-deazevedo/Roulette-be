CREATE TABLE IF NOT EXISTS restaurants(
    restaurant_id serial PRIMARY KEY,
    naam TEXT NOT NULL,
    picked NUMERIC DEFAULT 0,
    webUrl TEXT
);

CREATE TABLE IF NOT EXISTS users(
    user_id serial PRIMARY KEY,
    voornaam TEXT NOT NULL,
    achternaam TEXT NOT NULL
);


INSERT INTO restaurants(naam,webUrl) 
VALUES ('Subway', 'https://restaurants.subway.com/nl/nederland/heerlen/saroleastraat-61?utm_source=yxt-goog&utm_medium=local&utm_term=acq&utm_content=35154&utm_campaign=evergreen-2020'), 
('KFC', 'https://www.kfc.nl/'), 
('McDonalds', 'https://mcdonaldsrestaurant.nl/heerlen-woonboulevard');

INSERT INTO users(voornaam, achternaam)
VALUES ('Alain', 'Stoffels'), ('Bas', 'Verheule'), ('Ivan', 'de Azevedo'), ('Jesse', 'Kroon'), ('Nino', 'Theunissen'), 
('Ruud', 'Thissen'), ('Tom', 'Bessems'), ('Xander', 'den Heijer');