-- name: CreateStudent :exec
INSERT INTO students(email, name, department)
VALUES (
    ?,
    ?,
    ?
);

-- name: GetStudentByID :one
SELECT * FROM students
WHERE id = ?;

-- name: GetStudentByEmail :one
SELECT * FROM students
WHERE email = ?;

-- name: GetAllStudents :many
SELECT * FROM students;

-- name: UpdateStudent :exec
UPDATE students
SET 
    name = ?,
    department = ?
WHERE id = ?;

-- name: DeleteStudent :exec
DELETE FROM students WHERE id = ?;