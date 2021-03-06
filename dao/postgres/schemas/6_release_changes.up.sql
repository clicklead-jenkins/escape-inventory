ALTER TABLE release ADD COLUMN processed_dependencies BOOLEAN DEFAULT 'false';
ALTER TABLE release ADD COLUMN downloads INTEGER DEFAULT '0';
ALTER TABLE release ADD COLUMN uploaded_by VARCHAR(64) DEFAULT '';
ALTER TABLE release ADD COLUMN uploaded_at INTEGER DEFAULT '0';

ALTER TABLE application ADD COLUMN uploaded_by VARCHAR(64) DEFAULT '';
ALTER TABLE application ADD COLUMN uploaded_at INTEGER DEFAULT '0';

ALTER TABLE package ADD COLUMN uploaded_by VARCHAR(64) DEFAULT '';
ALTER TABLE package ADD COLUMN uploaded_at INTEGER DEFAULT '0';
ALTER TABLE package ADD COLUMN filesize INTEGER DEFAULT '-1';
