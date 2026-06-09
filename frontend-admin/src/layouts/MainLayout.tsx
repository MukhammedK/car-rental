import { Layout, Menu } from "antd";
import { Link, Outlet } from "react-router-dom";

const { Sider, Content } = Layout;

export default function MainLayout() {
    return (
        <Layout style={{ minHeight: "100vh" }}>
            <Sider>
                <Menu
                    theme="dark"
                    mode="inline"
                    items={[
                        {
                            key: "dashboard",
                            label: <Link to="/dashboard">Dashboard</Link>,
                        },
                        {
                            key: "cars",
                            label: <Link to="/cars">Cars</Link>,
                        },
                        {
                            key: "customers",
                            label: <Link to="/customers">Customers</Link>,
                        },
                        {
                            key: "bookings",
                            label: <Link to="/bookings">Bookings</Link>,
                        },
                    ]}
                />
            </Sider>

            <Content style={{ padding: 24 }}>
                <Outlet />
            </Content>
        </Layout>
    );
}