-- +goose Up
CREATE TABLE attendance(
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    student_id VARCHAR(36) NOT NULL,
    date DATE DEFAULT (CURRENT_DATE()),
    status ENUM('PRESENT', 'ABSENT') NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    UNIQUE (student_id, date)
);

-- +goose Down
DROP TABLE attendance;