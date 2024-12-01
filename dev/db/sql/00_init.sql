-- Create "nodes" table
CREATE TABLE bom (
  `id` text NOT NULL,
  `name` text NOT NULL,
  `version` text NOT NULL,
  `purl` text NOT NULL,
  PRIMARY KEY (`id`)
);