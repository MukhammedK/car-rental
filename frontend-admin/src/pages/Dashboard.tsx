import { useEffect, useState } from "react";
import api from "../api/client";
import { Card, Row, Col, Button} from "antd";
import { useNavigate } from "react-router-dom";

export default function Dashboard() {
    const [stats, setStats] = useState<any>(null);
    const navigate = useNavigate();
    const logout = () => {

        localStorage.removeItem(
            "token"
        );

        navigate("/");
    };


    useEffect(() => {
        const loadStats = async () => {
            try {
                const res = await api.get("/dashboard");
                setStats(res.data);
            } catch (err) {
                console.error(err);
            }
        };

        loadStats();
    }, []);

    if (!stats) {
        return <h2>Loading...</h2>;
    }

    return (
        <>
            <div
                style={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    marginBottom: 20,
                }}
            >
                <h1>Dashboard</h1>

                <Button
                    danger
                    onClick={logout}
                >
                    Logout
                </Button>
            </div>

            <Row gutter={16}>
                <Col span={6}>
                    <Card title="Cars Total">
                        <h2>{stats?.cars_total}</h2>
                    </Card>
                </Col>

                <Col span={6}>
                    <Card title="Cars Available">
                        <h2>{stats?.cars_available}</h2>
                    </Card>
                </Col>

                <Col span={6}>
                    <Card title="Bookings New">
                        <h2>{stats?.bookings_new}</h2>
                    </Card>
                </Col>

                <Col span={6}>
                    <Card title="Bookings Active">
                        <h2>{stats?.bookings_active}</h2>
                    </Card>
                </Col>
            </Row>

            <Row style={{ marginTop: 16 }}>
                <Col span={24}>
                    <Card title="Revenue">
                        <h1>{stats?.revenue} ₸</h1>
                    </Card>
                </Col>
            </Row>

        </>

    );
}