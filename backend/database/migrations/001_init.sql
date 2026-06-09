CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       full_name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       role VARCHAR(50) NOT NULL DEFAULT 'manager',
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE customers (
                           id SERIAL PRIMARY KEY,
                           full_name VARCHAR(255) NOT NULL,
                           phone VARCHAR(50) NOT NULL,
                           email VARCHAR(255),
                           iin VARCHAR(12),
                           created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE cars (
                      id SERIAL PRIMARY KEY,
                      brand VARCHAR(100) NOT NULL,
                      model VARCHAR(100) NOT NULL,
                      year INTEGER,
                      license_plate VARCHAR(50) UNIQUE,
                      color VARCHAR(50),
                      transmission VARCHAR(50),
                      fuel_type VARCHAR(50),
                      daily_price NUMERIC(10,2) NOT NULL,
                      status VARCHAR(50) DEFAULT 'available',
                      description TEXT,
                      created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE bookings (
                          id SERIAL PRIMARY KEY,

                          customer_id INTEGER REFERENCES customers(id),

                          car_id INTEGER REFERENCES cars(id),

                          start_date DATE NOT NULL,
                          end_date DATE NOT NULL,

                          total_price NUMERIC(10,2),

                          status VARCHAR(50) DEFAULT 'new',

                          source VARCHAR(50) DEFAULT 'website',

                          comment TEXT,

                          created_at TIMESTAMP DEFAULT NOW(),
                          updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE payments (
                          id SERIAL PRIMARY KEY,

                          booking_id INTEGER REFERENCES bookings(id),

                          amount NUMERIC(10,2) NOT NULL,

                          payment_method VARCHAR(50),

                          status VARCHAR(50),

                          payment_date TIMESTAMP DEFAULT NOW()
);

