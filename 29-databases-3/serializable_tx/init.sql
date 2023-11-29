CREATE TABLE rooms (
   room_id INT PRIMARY KEY,
   room_name VARCHAR(255) NOT NULL,
   is_available BOOLEAN NOT NULL
);

CREATE TABLE reservations (
      reservation_id SERIAL PRIMARY KEY,
      room_id INT NOT NULL,
      guest_name VARCHAR(255) NOT NULL,
      check_in_date DATE NOT NULL,
      check_out_date DATE NOT NULL,
      CONSTRAINT fk_room_id FOREIGN KEY (room_id) REFERENCES rooms (room_id)
);

-- Вставляем начальные данные о номерах
INSERT INTO rooms (room_id, room_name, is_available)
VALUES (1, 'Standard Room', true);

INSERT INTO rooms (room_id, room_name, is_available)
VALUES (2, 'Deluxe Suite', true);