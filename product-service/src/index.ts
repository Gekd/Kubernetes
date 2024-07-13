import express, { Request, Response } from 'express';
import productRoutes from './routes/productRoutes';

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());

app.use('/products', productRoutes);

app.get('/', (req: Request, res: Response) => {
  res.send('Product Service');
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
