CREATE TABLE Student_Class (
    id SERIAL PRIMARY KEY,
    class_id INT NOT NULL,
    student_id INT NOT NULL,
    CONSTRAINT fk_student_class_class
        FOREIGN KEY (class_id)
        REFERENCES Class(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_student_class_student
        FOREIGN KEY (student_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE
);
