CREATE TABLE IF NOT EXISTS categories
(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  m_category_class_id BIGINT UNSIGNED NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX m_category_class_index (m_category_class_id),
  FOREIGN KEY (m_category_class_id) REFERENCES m_category_classes (id)
);