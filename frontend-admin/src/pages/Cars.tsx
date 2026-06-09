import { useEffect, useState } from "react";
import api from "../api/client";
import { Table, Button, Modal, Form, Input,Popconfirm, message} from "antd";
import {useNavigate} from "react-router-dom";


export default function Cars() {
    const [cars, setCars] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [editingCar, setEditingCar] = useState<any>(null);
    const [form] = Form.useForm();
    const navigate = useNavigate();

    useEffect(() => {
        const loadCars = async () => {
            const res = await api.get("/cars");
            setCars(res.data);
        };

        loadCars();
    }, []);
    const logout = () => {

        localStorage.removeItem(
            "token"
        );

        navigate("/");
    };
    const handleEdit = (car: any) => {
        setEditingCar(car);

        form.setFieldsValue({
            license_plate: car.license_plate,
            brand: car.brand,
            model: car.model,
            year: car.year,
            daily_price: car.daily_price,
        });

        setIsModalOpen(true);
        message.success("Car updated");
    };

    const handleSave = async () => {
        const values = await form.validateFields();

        if (editingCar) {

            await api.put(`/cars/${editingCar.id}`, {
                license_plate: values.license_plate,
                brand: values.brand,
                model: values.model,
                year: Number(values.year),
                daily_price: Number(values.daily_price),
                status: editingCar.status,
            });

        } else {

            await api.post("/cars", {
                license_plate: values.license_plate,
                brand: values.brand,
                model: values.model,
                year: Number(values.year),
                daily_price: Number(values.daily_price),
                status: "available",
            });

        }

        const res = await api.get("/cars");
        setCars(res.data);

        form.resetFields();

        setEditingCar(null);

        setIsModalOpen(false);
        message.success("Car created");
    };


    const columns = [
        {
            title: "№",
            render: (_: any, __: any, index: number) => index + 1,
        },
        {
            title: "ID",
            dataIndex: "id",
        },
        {
            title: "Brand",
            dataIndex: "brand",
        },
        {
            title: "Model",
            dataIndex: "model",
        },
        {
            title: "Year",
            dataIndex: "year",
        },
        {
            title: "Price",
            dataIndex: "daily_price",
        },
        {
            title: "Status",
            dataIndex: "status",
        },
        {
            title: "Actions",
            render: (_: any, record: any) => (
                <>
                    <Button
                        onClick={() => handleEdit(record)}
                    >
                        Edit
                    </Button>

                    <Popconfirm
                        title="Delete car?"
                        onConfirm={() => handleDelete(record.id)}
                    >
                        <Button danger>
                            Delete
                        </Button>
                    </Popconfirm>
                </>
            ),
        }
    ];

    const handleDelete = async (id: number) => {
        try {
            await api.delete(`/cars/${id}`);

            const res = await api.get("/cars");
            setCars(res.data);
        } catch (err) {
            console.error(err);
        }
        message.success("Car deleted");
    };

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
                <h1>Cars</h1>

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
                    setEditingCar(null);

                    form.resetFields();

                    setIsModalOpen(true);
                }}
            >
                Add Car
            </Button>

            <Table
                rowKey="id"
                columns={columns}
                dataSource={cars}
            />
            <Modal
                title="Add Car"
                open={isModalOpen}
                onCancel={() => setIsModalOpen(false)}
                footer={null}
            >
                <Form
                    form={form}
                    layout="vertical"
                >
                    <Form.Item
                        label="License Plate"
                        name="license_plate"
                        rules={[
                            {
                                required: true,
                                message: "Enter license plate",
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="Brand"
                        name="brand"
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Model"
                        name="model"
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Year"
                        name="year"
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Price"
                        name="daily_price"
                    >
                        <Input />
                    </Form.Item>

                    <Button
                        type="primary"
                        onClick={handleSave}
                    >
                        Save
                    </Button>
                </Form>
            </Modal>
        </>
    );
}