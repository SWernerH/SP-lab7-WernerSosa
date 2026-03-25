-- -----------------------------------------------------------------------------
-- 1.  (Optional) Create database
-- -----------------------------------------------------------------------------
-- CREATE DATABASE university;

-- -----------------------------------------------------------------------------
-- 2.  STUDENTS TABLE
-- -----------------------------------------------------------------------------

DROP TABLE IF EXISTS students;

CREATE TABLE students (
    id         BIGSERIAL    PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    programme  TEXT         NOT NULL,
    year       SMALLINT     NOT NULL CHECK (year BETWEEN 1 AND 4),
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_students_name ON students (name);

-- -----------------------------------------------------------------------------
-- 3.  COURSES TABLES
-- -----------------------------------------------------------------------------

DROP TABLE IF EXISTS course_instructors;
DROP TABLE IF EXISTS courses;

CREATE TABLE courses (
    code     TEXT PRIMARY KEY,
    title    TEXT NOT NULL,
    credits  INT  NOT NULL,
    enrolled INT  DEFAULT 0
);

CREATE TABLE course_instructors (
    id          SERIAL PRIMARY KEY,
    course_code TEXT REFERENCES courses(code) ON DELETE CASCADE,
    instructor  TEXT NOT NULL
);

-- -----------------------------------------------------------------------------
-- 4.  SEED DATA (STUDENTS)
-- -----------------------------------------------------------------------------

INSERT INTO students (name, programme, year) VALUES
('Eve Castillo',   'BSc Computer Science',    2),
('Marco Tillett',  'BSc Computer Science',    3),
('Aisha Gentle',   'BSc Information Systems', 1),
('Raj Palacio',    'BSc Computer Science',    4);

-- -----------------------------------------------------------------------------
-- 5.  SEED DATA (COURSES)
-- -----------------------------------------------------------------------------

INSERT INTO courses (code, title, credits, enrolled) VALUES
('CMPS2212', 'GUI Programming', 3, 28),
('CMPS3412', 'Database Systems', 3, 22);

INSERT INTO course_instructors (course_code, instructor) VALUES
('CMPS2212', 'Boss'),
('CMPS3412', 'Dr. Ramos');

-- -----------------------------------------------------------------------------
-- 6.  VERIFY
-- -----------------------------------------------------------------------------

SELECT * FROM students;
SELECT * FROM courses;
SELECT * FROM course_instructors;