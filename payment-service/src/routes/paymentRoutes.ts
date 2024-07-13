import { Router } from 'express';
import { createPayment } from '../controllers/paymentController';

const router = Router();

router.post('/', createPayment);

export default router;
