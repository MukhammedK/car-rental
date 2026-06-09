import { createBrowserRouter } from "react-router-dom";

import Booking from "../pages/Booking.tsx";


export const router = createBrowserRouter([
    {
        path: "/",
        element: <Booking />,
    },




]);