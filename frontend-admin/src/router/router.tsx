import { createBrowserRouter } from "react-router-dom";

import Login from "../pages/Login";
import Dashboard from "../pages/Dashboard";
import Cars from "../pages/Cars";
import Customers from "../pages/Customers";
import Bookings from "../pages/Bookings";
import MainLayout from "../layouts/MainLayout";
import Register from "../pages/Register.tsx";
import ProtectedRoute from "../components/ProtectedRoute";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <Login />,
    },
    {
        path: "/register",
        element: <Register />,
    },
    {
        element: <MainLayout />,
        children: [

            {
                path: "/dashboard",
                element: (
                    <ProtectedRoute>
                        <Dashboard />
                    </ProtectedRoute>
                ),
            },
            {
                path: "/cars",
                element: (
                    <ProtectedRoute>
                        <Cars />
                    </ProtectedRoute>
                ),
            },
            {
                path: "/customers",
                element: (
                    <ProtectedRoute>
                        <Customers />
                    </ProtectedRoute>
                ),
            },
            {
                path: "/bookings",
                element: (
                    <ProtectedRoute>
                        <Bookings />
                    </ProtectedRoute>
                ),
            }

        ],
    },
]);