import type { AppProps } from "next/app";
import { Roboto } from "next/font/google";

import "@/styles/globals.css";
import Layout from "@/components/layout";

const roboto = Roboto({
  weight: "400",
  subsets: ["latin"],
});

export default function App({ Component, pageProps }: AppProps) {
  return (
    <main className={roboto.className}>
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </main>
  );
}
