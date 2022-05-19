ALTER TABLE network_config ADD COLUMN chain_color VARCHAR NOT NULL DEFAULT "";
UPDATE network_config SET chain_color = "";
ALTER TABLE network_config ADD COLUMN short_name VARCHAR NOT NULL DEFAULT "";
UPDATE network_config SET short_name = "";
