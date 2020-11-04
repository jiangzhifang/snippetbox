CREATE TABLE snippets (
	id serial primary key,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	created timestamp(0) NOT NULL,
	expires timestamp(0) NOT NULL
);
-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);
