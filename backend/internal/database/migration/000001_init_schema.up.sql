-- Создание расширения для UUID, если оно еще не установлено
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создание типа перечисления для организации
CREATE TYPE organization_type AS ENUM ('IE', 'LLC', 'JSC');

-- Создание таблицы пользователя
CREATE TABLE employee (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          username VARCHAR(50) UNIQUE NOT NULL,
                          first_name VARCHAR(50),
                          last_name VARCHAR(50),
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы организации
CREATE TABLE organization (
                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              name VARCHAR(100) NOT NULL,
                              description TEXT,
                              type organization_type,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы ответственного за организацию
CREATE TABLE organization_responsible (
                                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                          organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
                                          user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);


-- Вставка тестовых данных в таблицу пользователей
INSERT INTO employee (id, username, first_name, last_name) VALUES
                                                               (uuid_generate_v4(), 'user1', 'John', 'Doe'),
                                                               (uuid_generate_v4(), 'user2', 'Jane', 'Smith');

-- Вставка тестовых данных в таблицу организаций
INSERT INTO organization (id, name, description, type) VALUES
                                                           (uuid_generate_v4(), 'Org1', 'Test Organization 1', 'LLC'),
                                                           (uuid_generate_v4(), 'Org2', 'Test Organization 2', 'JSC');

-- Вставка тестовых данных в таблицу ответственных за организацию
INSERT INTO organization_responsible (id, organization_id, user_id) VALUES
                                                                        (uuid_generate_v4(), (SELECT id FROM organization WHERE name = 'Org1'), (SELECT id FROM employee WHERE username = 'user1')),
                                                                        (uuid_generate_v4(), (SELECT id FROM organization WHERE name = 'Org2'), (SELECT id FROM employee WHERE username = 'user2'));

-- Вставка тестовых данных в таблицу тендеров
INSERT INTO tender (id, organization_id, created_by, status, version) VALUES
                                                                          (uuid_generate_v4(), (SELECT id FROM organization WHERE name = 'Org1'), (SELECT id FROM employee WHERE username = 'user1'), 'CREATED', 1),
                                                                          (uuid_generate_v4(), (SELECT id FROM organization WHERE name = 'Org2'), (SELECT id FROM employee WHERE username = 'user2'), 'PUBLISHED', 1);

