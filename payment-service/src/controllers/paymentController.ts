import { Request, Response } from 'express';
import { Payment, CreatePayment } from '../models/paymentModel';

let payments: Payment[] = [];

export const createPayment = (req: Request, res: Response) => {
  const newPayment: CreatePayment = req.body;
  const id = payments.length ? payments[payments.length - 1].id + 1 : 1;
  const payment: Payment = { id, ...newPayment };
  payments.push(payment);
  res.status(201).json(payment);
};