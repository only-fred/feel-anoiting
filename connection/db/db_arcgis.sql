createdb db_arcgis

CREATE TABLE tbl_country_region(
  id SERIAL PRIMARY KEY,
  country_region VARCHAR(100) UNIQUE
);

CREATE TABLE tbl_last_update(
  id SERIAL PRIMARY KEY,
  last_update DECIMAL UNIQUE
);

CREATE TABLE tbl_attributes(
  id SERIAL PRIMARY KEY,
  objectid INT,
  province_state VARCHAR(100),
  id_country_region INT,
  id_last_update INT,
  confirmed INT,
  recovered INT,
  deaths INT,
  active INT,

  FOREIGN KEY (id_country_region)
  REFERENCES tbl_country_region(id),

  FOREIGN KEY (id_last_update)
  REFERENCES tbl_last_update(id)
);

INSERT INTO tbl_country_region (country_region)
  VALUES ('Brazil');
INSERT INTO tbl_last_update (last_update)
  VALUES (1594132439000);
INSERT INTO tbl_attributes (objectid, province_state, id_country_region, id_last_update, confirmed, recovered, deaths, active)
  VALUES (85, 'Ceara', 1, 1, 122477, 97470, 6481, 18526);

SELECT att.objectid, cr.country_region, lu.last_update, 
       att.confirmed, att.recovered, att.deaths, att.active
FROM tbl_attributes AS att
JOIN tbl_country_region AS cr
ON cr.id = att.id_country_region
JOIN tbl_last_update AS lu
ON lu.id = att.id_last_update;