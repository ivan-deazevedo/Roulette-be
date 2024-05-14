CREATE TABLE IF NOT EXISTS restaurants(
    restaurant_id serial PRIMARY KEY,
    naam TEXT NOT NULL,
    picked NUMERIC DEFAULT 0
);


INSERT INTO restaurants(naam) VALUES ('Subway'), ('KFC'), ('McDonalds') 