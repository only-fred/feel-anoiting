createdb db_arcgis

CREATE TABLE tbl_attributes(
  id SERIAL PRIMARY KEY,
  province_state VARCHAR(100),
  country_region VARCHAR(100),
  last_update VARCHAR(20),
  confirmed INT,
  recovered INT,
  deaths INT,
  active INT
);

INSERT INTO tbl_attributes (province_state, country_region, last_update, confirmed, recovered, deaths, active)
  VALUES ('Ceara', 'Brazil', '03/03/03', 122477, 97470, 6481, 18526);
