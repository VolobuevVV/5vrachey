INSERT INTO departments (id, name, color, is_active) VALUES
('gynecology', 'Гинекология', '#fce4ec', true),
('urology', 'Урология', '#e3f2fd', true),
('neurology', 'Неврология', '#fff3e0', true),
('cardiology', 'Кардиология', '#ffebee', true),
('gastroenterology', 'Гастроэнтерология', '#f3e5f5', true),
('dermatology', 'Дерматология', '#e8f5e9', true),
('pediatrics', 'Педиатрия', '#fff8e1', true),
('therapy', 'Терапия', '#e0f7fa', true),
('diagnostics', 'Диагностика', '#f5f5f5', true),
('rehabilitation', 'Реабилитация', '#e8eaf6', true),
('endocrinology', 'Эндокринология', '#f1f8e9', true),
('psychotherapy', 'Психотерапия', '#f3e5f5', true)
ON CONFLICT (id) DO NOTHING;