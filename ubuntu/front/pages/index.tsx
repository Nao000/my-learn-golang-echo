import Head from 'next/head'

// @ を使ってパスを指定するとVScodeで定義にいけなかったので修正
import styles from '../styles/Home.module.css'

import Employee from '@/interfaces/employee'

import { loadEmployees } from '@/lib/load-posts';

export async function getStaticProps() {

  const employees =await loadEmployees();

  return { props: { employees } };
}

type EmployeeArr = {
  employees: Employee[]
}

export default function Home({employees} :EmployeeArr) {
   return (
    <>
      <Head>
        <title>Create Next App</title>
      </Head>
      <main className={styles["main"]}>
        <ul className={styles["card-ul"]}>
          {employees.map(({ emp_no, first_name, last_name, hire_date}) => (
            <li className={styles["card-li"]} key={emp_no}>
              <p>emp_no: {emp_no}</p>
              <p>first_name: {first_name}</p>
              <p>last_name: {last_name}</p>
              <p>hire_date: {hire_date}</p>
            </li>
          ))}
        </ul>


      </main>
    </>
  )
}
