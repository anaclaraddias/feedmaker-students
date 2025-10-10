CREATE TABLE Class (
    id SERIAL PRIMARY KEY,
    teacher_id INT NOT NULL,
    student_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255),
    CONSTRAINT fk_class_student
        FOREIGN KEY (student_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_class_teacher
        FOREIGN KEY (teacher_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE
);