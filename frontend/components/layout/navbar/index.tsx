import Button from "@/components/button";
import Image from "next/image";
import Link from "next/link";

type Props = {
  toogleDrawer: () => void;
};

function Navbar({ toogleDrawer }: Props) {
  const onClickHamburger = () => {
    toogleDrawer();
  };
  return (
    <nav className="fixed flex items-center justify-between w-screen px-5 py-3 bg-white shadow">
      <Link href="/">
        <div className="flex items-end gap-2">
          <Image src="svg/logo-sm.svg" height={50} width={50} alt="logo" />
          <p className="hidden text-secondary-500 lg:block">
            Lo importante es disfrutar el camino
          </p>
        </div>
      </Link>
      <section className="flex gap-5">
        <div className="hidden md:block w-36">
          <Button
            onClick={() => console.log("button clicked")}
            variant="outlined"
          >
            Crear cuenta
          </Button>
        </div>
        <div className="hidden md:block w-36">
          <Button
            onClick={() => console.log("button clicked")}
            variant="outlined"
          >
            Iniciar sesión
          </Button>
        </div>
        <button onClick={onClickHamburger} className="md:hidden">
          <i className="fa-solid fa-bars"></i>
        </button>
      </section>
    </nav>
  );
}

export default Navbar;
