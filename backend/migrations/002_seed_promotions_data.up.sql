INSERT INTO promotions (text, is_active)
VALUES ('Скидка 20% на первое посещение', true)
ON CONFLICT (id) DO NOTHING;