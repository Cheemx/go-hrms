-- name: MarkAttendance :exec
INSERT INTO attendance(student_id, status)
VALUES (
    ?,
    ?
);

-- name: GetAttendanceByStudentID :many
SELECT 
    students.name AS name,
    students.email AS email,
    attendance.status AS status,
    attendance.date AS date
FROM attendance 
JOIN students
ON students.id = attendance.student_id
WHERE attendance.student_id = ?;

-- name: GetAttendanceSummary :one
SELECT
    SUM(CASE WHEN status = 'PRESENT' THEN 1 ELSE 0 END) AS present_days,
    COUNT(*) AS total_days
FROM attendance
WHERE student_id = ?;