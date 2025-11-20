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

-- name: GetWeeklyAttendanceSummary :many
SELECT
    a.student_id AS student_id,
    s.name AS name,
    s.email AS email,
    CAST(SUM(CASE WHEN a.status = 'PRESENT' THEN 1 ELSE 0 END) AS UNSIGNED) AS present_days,
    COUNT(*) AS total_days
FROM attendance a
JOIN students s 
ON s.id = a.student_id
WHERE a.date >= DATE_SUB(CURDATE(), INTERVAL 6 DAY)
AND a.date <= CURDATE()
GROUP BY a.student_id, s.name, s.email;

-- name: GetMonthlyAttendanceSummary :many
SELECT
    a.student_id AS student_id,
    s.name AS name,
    s.email AS email,
    CAST(SUM(CASE WHEN a.status = 'PRESENT' THEN 1 ELSE 0 END) AS UNSIGNED) AS present_days,
    COUNT(*) AS total_days
FROM attendance a
JOIN students s 
ON s.id = a.student_id
WHERE a.date >= DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 1 MONTH), '%Y-%m-01')
AND a.date <  DATE_FORMAT(CURDATE(), '%Y-%m-01')
GROUP BY a.student_id, s.name, s.email;