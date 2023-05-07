import { Html, Head, Main, NextScript } from 'next/document'
import Footer from "@/components/footer"

export default function Document() {
  return (
    <Html lang="ja">
      <Head />
      <body>
        <Main />
        <NextScript />
        <Footer exampleArg="コンポーネントに引数を渡す" />
      </body>
    </Html>
  )
}
