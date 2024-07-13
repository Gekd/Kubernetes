import express, { Request, Response } from 'express';
import paymentRoutes from './routes/paymentRoutes';

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());

app.use('/payments', paymentRoutes);

app.get('/', (req: Request, res: Response) => {
  res.send('Payment Service');
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
