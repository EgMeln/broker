CREATE TABLE IF NOT EXISTS positions(
    ID         UUID PRIMARY KEY,
    price_open  float,
    is_bay     BOOLEAN,
    symbol     VARCHAR(60),
    price_close float
);