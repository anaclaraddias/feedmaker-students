CREATE TABLE Feedback (
    id SERIAL PRIMARY KEY,
    score INT NOT NULL,
    body TEXT NOT NULL,
    student_id INT NOT NULL,
    teacher_id INT NOT NULL,
    CONSTRAINT fk_feedback_student
        FOREIGN KEY (student_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_feedback_teacher
        FOREIGN KEY (teacher_id)
        REFERENCES "user"(id)
        ON DELETE CASCADE
);
