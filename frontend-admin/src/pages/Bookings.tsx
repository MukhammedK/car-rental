import { useEffect, useState } from "react";
import api from "../api/client";
import {
    Table,
    Button,
    Modal,
    Form,
    Input,
    Select,
    DatePicker,
    message,
} from "antd";
import {useNavigate} from "react-router-dom";

export default function Bookings() {
    const [bookings, setBookings] = useState([]);
    const [customers, setCustomers] = useState<any[]>([]);
    const [cars, setCars] = useState<any[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const navigate = useNavigate();


    const [form] = Form.useForm();

    useEffect(() => {
        loadBookings();
        loadCustomers();
        loadCars();
    }, []);

    const loadBookings = async () => {
        const res = await api.get("/bookings");
        setBookings(res.data);
    };

    const loadCustomers = async () => {
        const res = await api.get("/customers");
        setCustomers(res.data);
    };

    const loadCars = async () => {
        const res = await api.get("/cars");
        setCars(res.data);
    };
    const updateStatus = async (
        id: number,
        status: string,
    ) => {
        try {

            await api.patch(
                `/bookings/${id}/status`,
                {
                    status,
                }
            );

            const res = await api.get(
                "/bookings"
            );

            setBookings(res.data);

            message.success(
                "Status updated"
            );

        } catch (err) {

            console.error(err);

            message.error(
                "Failed to update status"
            );
        }

    };
    const logout = () => {

        localStorage.removeItem(
            "token"
        );

        navigate("/");
    };

    const handleSaveBooking = async () => {
        try {
            const values = await form.validateFields();

            await api.post("/bookings", {
                customer_id: values.customer_id,
                car_id: values.car_id,
                start_date: values.start_date,
                end_date: values.end_date,
                total_price: Number(values.total_price),
                comment: values.comment,
            });

            await loadBookings();

            form.resetFields();
            setIsModalOpen(false);

            message.success("Booking created");
        } catch (err: any) {

            const errorMessage =
                err?.response?.data?.error;

            message.error(
                errorMessage || "Booking failed"
            );
        }
    };

    const columns = [
        {
            title: "№",
            render: (_: any, __: any, index: number) => index + 1,
        },
        {
            title: "Customer",
            render: (_: any, record: any) => {

                const customer = customers.find(
                    (c: any) =>
                        c.id === record.customer_id
                );

                return customer
                    ? customer.full_name
                    : record.customer_id;
            },
        },
        {
            title: "Car",
            render: (_: any, record: any) => {

                const car = cars.find(
                    (c: any) =>
                        c.id === record.car_id
                );

                return car
                    ? `${car.brand} ${car.model}`
                    : record.car_id;
            },
        },
        {
            title: "Start",
            dataIndex: "start_date",
        },
        {
            title: "End",
            dataIndex: "end_date",
        },
        {
            title: "Price",
            dataIndex: "total_price",
        },
        {
            title: "Status",
            dataIndex: "status",
        },
        {
            title: "Status",
            render: (_: any, record: any) => {

                const isLocked =
                    record.status === "completed" ||
                    record.status === "cancelled";

                return (
                    <Select
                        value={record.status}
                        disabled={isLocked}
                        style={{ width: 150 }}
                        onChange={(value) =>
                            updateStatus(
                                record.id,
                                value
                            )
                        }
                        options={[
                            {
                                value: "new",
                                label: "New",
                            },
                            {
                                value: "approved",
                                label: "Approved",
                            },
                            {
                                value: "completed",
                                label: "Completed",
                            },
                            {
                                value: "cancelled",
                                label: "Cancelled",
                            },
                        ]}
                    />
                );
            },
        }
    ];

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
                <h1>Bookings</h1>

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
                onClick={() => setIsModalOpen(true)}
            >
                Add Booking
            </Button>

            <Table
                rowKey="id"
                columns={columns}
                dataSource={bookings}
            />

            <Modal
                title="Add Booking"
                open={isModalOpen}
                onCancel={() => setIsModalOpen(false)}
                footer={null}
            >
                <Form
                    form={form}
                    layout="vertical"
                >
                    <Form.Item
                        label="Customer"
                        name="customer_id"
                    >
                        <Select>
                            {customers.map((customer: any) => (
                                <Select.Option
                                    key={customer.id}
                                    value={customer.id}
                                >
                                    {customer.full_name}
                                </Select.Option>
                            ))}
                        </Select>
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
                        <DatePicker style={{ width: "100%" }} />
                    </Form.Item>

                    <Form.Item
                        label="End Date"
                        name="end_date"
                    >
                        <DatePicker style={{ width: "100%" }} />
                    </Form.Item>

                    <Form.Item
                        label="Total Price"
                        name="total_price"
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Comment"
                        name="comment"
                    >
                        <Input />
                    </Form.Item>

                    <Button
                        type="primary"
                        onClick={handleSaveBooking}
                    >
                        Save
                    </Button>
                </Form>


            </Modal>


        </>
    );
}