import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Card, Input, Button } from "antd";
import api from "../api/client";

export default function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate();

    const login = async () => {
        try {
            const res = await api.post("/auth/login", {
                email,
                password,
            });

            localStorage.setItem("token", res.data.token);

            navigate("/dashboard");
        } catch {
            alert("Login failed");
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
                <Input
                    placeholder="Email"
                    style={{ marginBottom: 12 }}
                    onChange={(e) =>
                        setEmail(e.target.value)
                    }
                />

                <Input.Password
                    placeholder="Password"
                    style={{ marginBottom: 12 }}
                    onChange={(e) =>
                        setPassword(e.target.value)
                    }
                />

                <Button
                    type="primary"
                    block
                    onClick={login}
                >
                    Login
                </Button>

                <div
                    style={{
                        marginTop: 16,
                        textAlign: "center",
                    }}
                >
                    <Link to="/register">
                        Register
                    </Link>
                </div>
            </Card>
        </div>
    );
}