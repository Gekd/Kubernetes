import express, { Request, Response } from 'express';
import orderRoutes from './routes/orderRoutes';

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());

app.use('/orders', orderRoutes);

app.get('/', (req: Request, res: Response) => {
  res.send('Order Service');
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
