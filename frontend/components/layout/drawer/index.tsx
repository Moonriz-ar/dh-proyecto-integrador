import Link from "next/link";
import Social from "../social";

type Props = {
  toogleDrawer: () => void;
};

function Drawer({ toogleDrawer }: Props) {
  const onClickLink = () => {
    toogleDrawer();
  };

  const onClose = () => {
    toogleDrawer();
  };

  return (
    <div className="fixed inset-x-0 z-20 flex flex-col w-full h-screen bg-white md:hidden">
      <header className="relative flex items-end justify-end p-5 text-white bg-primary grow-0 shrink basis-1/5">
        <span
          className="absolute text-2xl inset-5 w-min h-min"
          onClick={onClose}
        >
          x
        </span>
        <h2>MENÚ</h2>
      </header>
      <section className="flex flex-col gap-2 p-5 text-right text-secondary grow shrink-0 basis-0">
        <Link href="" onClick={onClickLink}>
          <h3>Crear cuenta</h3>
        </Link>
        <hr />
        <Link href="" onClick={onClickLink}>
          <h3>Iniciar sesión</h3>
        </Link>
      </section>
      <section className="flex justify-end gap-3 p-5 text-secondary">
        <Social />
      </section>
    </div>
  );
}

export default Drawer;
