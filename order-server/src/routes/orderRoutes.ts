import { Router } from 'express';
import { createOrder, getOrderById, getAllOrders, updateOrder, deleteOrder } from '../controllers/orderController';

const router = Router();

router.post('/', createOrder);
router.get('/:id', getOrderById);
router.get('/', getAllOrders);
router.put('/:id', updateOrder);
router.delete('/:id', deleteOrder);

export default router;
