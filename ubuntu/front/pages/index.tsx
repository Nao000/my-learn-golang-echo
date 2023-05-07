import Head from 'next/head'
import styles from '@/styles/Home.module.css'
import Post from '@/interfaces/post'

import { loadPosts } from '@/lib/load-posts';

export async function getStaticProps() {
  const posts = await loadPosts();

  return { props: { posts } };
}

type Props = {
  posts: Post
}

export default function Home({posts}: Props) {

  return (
    <>
      <Head>
        <title>Create Next App</title>
      </Head>
      <main>
        Hello World
        <p>id: {posts.id}</p>
        <p>title: {posts.title}</p>
        <p>body: {posts.body}</p>
      </main>
    </>
  )
}
