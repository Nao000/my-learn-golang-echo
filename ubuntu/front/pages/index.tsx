import Head from 'next/head'
import styles from '@/styles/Home.module.css'

export default function Home() {
  return (
    <>
      <Head>
        <title>Create Next App</title>
      </Head>
      <main className={`${styles.main}`}>
        Hello World
      </main>
    </>
  )
}
