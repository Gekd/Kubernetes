import { Request, Response } from 'express';
import { Order, CreateOrder } from '../models/orderModel';

let orders: Order[] = [];

export const createOrder = (req: Request, res: Response) => {
  const newOrder: CreateOrder = req.body;
  const id = orders.length ? orders[orders.length - 1].id + 1 : 1;
  const order: Order = { id, ...newOrder };
  orders.push(order);
  res.status(201).json(order);
};

export const getOrderById = (req: Request, res: Response) => {
  const id = parseInt(req.params.id);
  const order = orders.find(o => o.id === id);
  if (order) {
    res.json(order);
  } else {
    res.status(404).send('Order not found');
  }
};

export const getAllOrders = (req: Request, res: Response) => {
  res.json(orders);
};

export const updateOrder = (req: Request, res: Response) => {
  const id = parseInt(req.params.id);
  const index = orders.findIndex(o => o.id === id);
  if (index !== -1) {
    const updatedOrder = { ...orders[index], ...req.body };
    orders[index] = updatedOrder;
    res.json(updatedOrder);
  } else {
    res.status(404).send('Order not found');
  }
};

export const deleteOrder = (req: Request, res: Response) => {
  const id = parseInt(req.params.id);
  const index = orders.findIndex(o => o.id === id);
  if (index !== -1) {
    orders.splice(index, 1);
    res.status(204).send();
  } else {
    res.status(404).send('Order not found');
  }
};
