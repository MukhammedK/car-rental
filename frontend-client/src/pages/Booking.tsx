import { useEffect, useState } from "react";
import { Form, Input, Button, Select, DatePicker, Card } from "antd";
import api from "../api/client";

export default function Booking() {
    const [form] = Form.useForm();
    const [cars, setCars] = useState<any[]>([]);

    useEffect(() => {
        const loadCars = async () => {
            const res = await api.get("/cars");

            console.log(res.data);

            setCars(res.data);
        };

        loadCars();
    }, []);
    const sendBooking = async () => {
        try {
            const values = await form.validateFields();

            const customerRes = await api.post(
                "/public/customers",
                {
                    full_name: values.full_name,
                    phone: values.phone,
                    email: values.email,
                    iin: values.iin,
                }
            );

            await api.post(
                "/public/bookings",
                {
                    customer_id: customerRes.data.id,
                    car_id: values.car_id,
                    start_date: values.start_date,
                    end_date: values.end_date,
                    comment: "Website booking",
                }
            );

            alert("Booking sent!");

            form.resetFields();

        } catch (err) {
            console.error(err);
            alert("Error sending booking");
        }
    };

    return (
        <div
            style={{
                minHeight: "100vh",
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                background: "#f5f5f5",
            }}
        >
            <Card
                title="Book a Car"
                style={{ width: 500 }}
            >
                <Form
                    form={form}
                    layout="vertical"
                >

                    <Form.Item
                        name="full_name"
                        rules={[
                            {
                                required: true,
                                message: "Enter full name",
                            },
                        ]}
                    >
                        <Input placeholder="Full Name" />
                    </Form.Item>

                    <Form.Item
                        label="IIN"
                        name="iin"
                        rules={[
                            {
                                required: true,
                                message: "Enter IIN",
                            },
                            {
                                pattern: /^\d{12}$/,
                                message: "IIN must contain exactly 12 digits",
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Phone"
                        name="phone"
                        rules={[
                            {
                                required: true,
                                message: "Enter phone number",
                            },
                            {
                                pattern: /^\+7\d{10}$/,
                                message: "Phone must be like +77001234567",
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        name="email"
                        rules={[
                            {
                                required: true,
                                message: "Enter email",
                            },
                            {
                                type: "email",
                                message: "Invalid email format",
                            },
                        ]}
                    >
                        <Input placeholder="Email" />
                    </Form.Item>

                    <Form.Item
                        label="Car"
                        name="car_id"
                    >
                        <Select>
                            {cars.map((car: any) => (
                                <Select.Option
                                    key={car.id}
                                    value={car.id}
                                >
                                    {car.brand} {car.model}
                                </Select.Option>
                            ))}
                        </Select>
                    </Form.Item>

                    <Form.Item
                        label="Start Date"
                        name="start_date"
                    >
                        <DatePicker
                            style={{ width: "100%" }}
                        />
                    </Form.Item>

                    <Form.Item
                        label="End Date"
                        name="end_date"
                    >
                        <DatePicker
                            style={{ width: "100%" }}
                        />
                    </Form.Item>

                    <Button
                        type="primary"
                        block
                        onClick={sendBooking}
                    >
                        Send Booking
                    </Button>

                </Form>
            </Card>
        </div>
    );
}