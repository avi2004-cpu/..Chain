import express from 'express';
import cors from 'cors';
import 'dotenv/config';

import identityRoutes from './routes/identity.js';
import assetRoutes from './routes/assets.js';
import accessRoutes from './routes/access.js';

const app = express();
app.use(cors());
app.use(express.json());

app.use('/identity', identityRoutes);
app.use('/assets', assetRoutes);
app.use('/access', accessRoutes);

const PORT = process.env.PORT || 4000;
app.listen(PORT, () => console.log(`ChainGuard backend running on port ${PORT}`));
