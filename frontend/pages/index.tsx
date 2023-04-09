import Button from "@/components/button";
import Layout from "@/components/layout";

export default function Home() {
  return (
    <Layout>
      <div>
        <h1 className="text-primary">hi</h1>
        <div className="w-36">
          <Button onClick={() => console.log("button clicked")}>Click</Button>
        </div>
      </div>
    </Layout>
  );
}
