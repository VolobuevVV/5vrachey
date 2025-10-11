INSERT INTO departments (id, name, full_name, color, is_active, position, head_doctor_id) VALUES
('gynecology', 'Гинекология', 'Гинекологическое отделение', '#fce4ec', true, 1, '283136-leonteva'),
('urology', 'Урология', 'Урологическое отделение', '#e3f2fd', true, 2, NULL),
('neurology', 'Неврология', 'Неврологическое отделение', '#fff3e0', true, 3, '167147-kokorin'),
('cardiology', 'Кардиология', 'Кардиологическое отделение', '#ffebee', true, 4, NULL),
('gastroenterology', 'Гастроэнтерология', 'Гастроэнтерологическое отделение', '#f3e5f5', true, 5, NULL),
('dermatology', 'Дерматология', 'Дерматологическое отделение', '#e8f5e9', true, 6, '301580-pashyan'),
('pediatrics', 'Педиатрия', 'Педиатрическое отделение', '#fff8e1', true, 7, NULL),
('therapy', 'Терапия', 'Терапевтическое отделение', '#e0f7fa', true, 8, NULL),
('diagnostics', 'Диагностика', 'Диагностическое отделение', '#f5f5f5', true, 9, NULL),
('rehabilitation', 'Реабилитация', 'Реабилитационное отделение', '#e8eaf6', true, 10, NULL),
('endocrinology', 'Эндокринология', 'Эндокринологическое отделение', '#f1f8e9', true, 11, NULL),
('psychotherapy', 'Психотерапия', 'Психотерапевтическое отделение', '#f3e5f5', true, 12, NULL)
ON CONFLICT (id) DO NOTHING;