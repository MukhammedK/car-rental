import { useEffect, useState } from "react";
import api from "../api/client";
import { Table, Button, Modal, Form, Input,Popconfirm, message} from "antd";
import {useNavigate} from "react-router-dom";

export default function Customers() {
    const [customers, setCustomers] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [editingCustomer, setEditingCustomer] = useState<any>(null);
    const [form] = Form.useForm();
    const navigate = useNavigate();


    useEffect(() => {
        const loadCustomers = async () => {
            const res = await api.get("/customers");
            setCustomers(res.data);
        };

        loadCustomers();
    }, []);
    const handleEdi = (customer: any) => {
        setEditingCustomer(customer);

        form.setFieldsValue({
            full_name: customer.full_name,
            phone: customer.phone,
            email: customer.email,
            iin: customer.iin,
        });

        setIsModalOpen(true);
        message.success("Customer updated");
    };

    const handleSav = async () => {
        const values = await form.validateFields();

        if (editingCustomer) {

            await api.put(`/customers/${editingCustomer.id}`, {
                full_name: values.full_name,
                phone: values.phone,
                email: values.email,
                iin: values.iin,
            });

        } else {

            await api.post("/customers", {
                full_name: values.full_name,
                phone: values.phone,
                email: values.email,
                iin: values.iin,
            });

        }

        const res = await api.get("/customers");
        setCustomers(res.data);

        form.resetFields();

        setEditingCustomer(null);

        setIsModalOpen(false);
        message.success("Customer created");
    };


    const columns = [
        {
            title: "№",
            render: (_: any, __: any, index: number) => index + 1,
        },
        {
            title: "Full Name",
            dataIndex: "full_name",
        },
        {
            title: "Phone",
            dataIndex: "phone",
        },
        {
            title: "Email",
            dataIndex: "email",
        },
        {
            title: "IIN",
            dataIndex: "iin",
        },
        {
            title: "Actions",
            render: (_: any, record: any) => (
                <>
                    <Button onClick={() => handleEdi(record)}>
                        Edit
                    </Button>

                    <Popconfirm
                        title="Delete customer?"
                        onConfirm={() => handleDelet(record.id)}
                    >
                        <Button danger>
                            Delete
                        </Button>
                    </Popconfirm>
                </>
            ),
        },
    ];
    const logout = () => {

        localStorage.removeItem(
            "token"
        );

        navigate("/");
    };

    const handleDelet = async (id: number) => {
        try {
            await api.delete(`/customers/${id}`);

            const res = await api.get("/customers");
            setCustomers(res.data);
        } catch (err) {
            console.error(err);
        }
        message.success("Customer deleted");
    };
    console.log(customers);

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
                <h1>Customers</h1>

                <Button
                    danger
                    onClick={logout}
                >
                    Logout
                </Button>
            </div>
            <Button
                type="primary"
                style={{ marginBottom: 16 }}
                onClick={() => {
                    setEditingCustomer(null);

                    form.resetFields();

                    setIsModalOpen(true);
                }}
            >
                Add Customer
            </Button>

            <Table
                rowKey="id"
                columns={columns}
                dataSource={customers}
            />
            <Modal
                title="Add Customer"
                open={isModalOpen}
                onCancel={() => setIsModalOpen(false)}
                footer={null}
            >
                <Form
                    form={form}
                    layout="vertical"
                >
                    <Form.Item
                        label="Full Name"
                        name="full_name"
                        rules={[
                            {
                                required: true,
                                message: "Enter full name",
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
                        label="Email"
                        name="email"
                    >
                        <Input />
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
                    <Button
                        type="primary"
                        onClick={handleSav}
                    >
                        Save
                    </Button>
                </Form>
            </Modal>
        </>
    );
}