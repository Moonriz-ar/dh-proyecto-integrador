import Social from "@/components/layout/social";

function Footer() {
  return (
    <footer className="flex justify-between w-screen px-5 py-3 item-center bg-primary text-ivory">
      <p className="text-sm">@2021 Digital Booking</p>
      <div className="hidden md:flex md:gap-5 md:items-end">
        <Social />
      </div>
    </footer>
  );
}

export default Footer;
