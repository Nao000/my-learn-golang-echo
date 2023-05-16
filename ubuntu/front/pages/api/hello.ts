// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  name: string,
  method: string,
}

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {

  if (req.method !== 'POST') {
    res.status(405).send({ name: 'NOT POST REQUEST', method: req.method, })

    return
  }

  res.status(200).json({ name: req.body.name, method: req.method })
}
