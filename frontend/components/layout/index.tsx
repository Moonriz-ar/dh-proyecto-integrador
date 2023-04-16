import { useState } from "react";

import Navbar from "@/components/layout/navbar";
import Footer from "@/components/layout/footer";
import Drawer from "@/components/layout/drawer";

type Props = {
  children: React.ReactElement;
};

function Layout({ children }: Props) {
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);

  const toogleDrawer = () => {
    setIsDrawerOpen((prevState) => !prevState);
  };

  return (
    <div className="relative flex flex-col h-full min-h-screen bg-ivory-500">
      <Navbar toogleDrawer={toogleDrawer} />
      <section className="px-5 pt-20 pb-2 grow shrink-0 basis-0">
        {children}
      </section>
      <Footer />

      {isDrawerOpen && <Drawer toogleDrawer={toogleDrawer} />}
    </div>
  );
}

export default Layout;
