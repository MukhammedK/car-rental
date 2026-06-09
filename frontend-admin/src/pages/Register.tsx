// import { useState } from "react";
import api from "../api/client";
import { Link } from "react-router-dom";
import { Card, Input, Button, message,Form } from "antd";
import { useNavigate } from "react-router-dom";

export default function Register() {


    // const [fullName] = useState("");
    // const [email] = useState("");
    // const [password] = useState("");
    const [form] = Form.useForm();
    const navigate = useNavigate();

    const register = async () => {
        try {
            const values = await form.validateFields();


            await api.post(
                "/auth/register",
                values
            );

            message.success(
                "Registration successful"
            );

            navigate("/dashboard");

        } catch {
            message.error(
                "Registration failed"
            );
        }
    };

    return (
        <div
            style={{
                height: "100vh",
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                background: "#f5f5f5",
            }}
        >
            <Card
                title="Car Rental CRM"
                style={{
                    width: 400,
                }}
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
                        name="password"
                        rules={[
                            {
                                required: true,
                                message: "Enter password",
                            },
                            {
                                min: 6,
                                message:
                                    "Password must contain at least 6 characters",
                            },
                        ]}
                    >
                        <Input.Password placeholder="Password" />
                    </Form.Item>

                    <Button
                        type="primary"
                        block
                        onClick={register}
                    >
                        Register
                    </Button>

                    <div
                        style={{
                            marginTop: 16,
                            textAlign: "center",
                        }}
                    >
                        <Link to="/">
                            Already have an account? Login
                        </Link>
                    </div>
                </Form>
            </Card>
        </div>
    );
}